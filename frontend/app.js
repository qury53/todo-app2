const API_BASE = "http://localhost:8080/api/v1"; 

// --- User state ---
let currentUserId = null;

const createUserForm = document.getElementById("create-user-form");
const currentUserInfo = document.getElementById("current-user-info");
const currentUserNameEl = document.getElementById("current-user-name");
const currentUserIdEl = document.getElementById("current-user-id");
const clearUserBtn = document.getElementById("clear-user-btn");

const createTaskForm = document.getElementById("create-task-form");
const tasksListEl = document.getElementById("tasks-list");
const loadTasksBtn = document.getElementById("load-tasks-btn");

// --- Helpers ---

async function apiRequest(path, options = {}) {
  const url = API_BASE + path;
  const res = await fetch(url, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
    },
  });

  let data = null;
  const text = await res.text();
  if (text) {
    try {
      data = JSON.parse(text);
    } catch {
      // ignore
    }
  }

  if (!res.ok) {
    const message = data?.message || data?.error || `HTTP ${res.status}`;
    throw new Error(message);
  }

  return data;
}

function saveUser(userId, fullName) {
  localStorage.setItem("todo_user_id", String(userId));
  localStorage.setItem("todo_user_name", fullName);
  currentUserId = userId;
  renderUserState();
}

function clearUser() {
  localStorage.removeItem("todo_user_id");
  localStorage.removeItem("todo_user_name");
  currentUserId = null;
  renderUserState();
}

function loadUserFromStorage() {
  const id = localStorage.getItem("todo_user_id");
  const name = localStorage.getItem("todo_user_name");
  if (id) {
    currentUserId = Number(id);
    currentUserNameEl.textContent = name || "";
    currentUserIdEl.textContent = id;
  }
}

function renderUserState() {
  if (currentUserId) {
    createUserForm.style.display = "none";
    currentUserInfo.style.display = "block";
    createTaskForm.style.display = "flex";
    loadTasksBtn.style.display = "inline-block";
    loadTasks();
  } else {
    createUserForm.style.display = "flex";
    currentUserInfo.style.display = "none";
    createTaskForm.style.display = "none";
    loadTasksBtn.style.display = "none";
    tasksListEl.innerHTML = "";
  }
}

// --- User actions ---

createUserForm.addEventListener("submit", async (e) => {
  e.preventDefault();

  const fullName = document.getElementById("user-full-name").value.trim();
  const email = document.getElementById("user-email").value.trim();
  const password = document.getElementById("user-password").value;
  const phone = document.getElementById("user-phone").value.trim();

  const body = {
    full_name: fullName,
    email,
    password,
  };

  if (phone) {
    body.phone_number = phone;
  }

  try {
    const res = await apiRequest("/users/", {
      method: "POST",
      body: JSON.stringify(body),
    });

    // res ~ CreateUserResponse
    saveUser(res.id, res.full_name || fullName);
    createUserForm.reset();
  } catch (err) {
    alert("Ошибка при создании пользователя: " + err.message);
  }
});

clearUserBtn.addEventListener("click", () => {
  if (confirm("Сменить пользователя? Текущие задачи останутся, но создаваться новые будут от нового пользователя.")) {
    clearUser();
  }
});

// --- Tasks actions ---

createTaskForm.addEventListener("submit", async (e) => {
  e.preventDefault();

  const title = document.getElementById("task-title").value.trim();
  const description = document.getElementById("task-description").value.trim();
  const completed = document.getElementById("task-completed").checked;

  const body = {
    author_id: currentUserId,
    title,
    description,
    completed,
  };

  try {
    await apiRequest("/tasks/", {
      method: "POST",
      body: JSON.stringify(body),
    });

    createTaskForm.reset();
    loadTasks();
  } catch (err) {
    alert("Ошибка при создании задачи: " + err.message);
  }
});

loadTasksBtn.addEventListener("click", loadTasks);

async function loadTasks() {
  tasksListEl.innerHTML = "Загрузка...";

  try {
    // GET /tasks/ -> GetTaskResponse (судя по схеме, один объект)
    // Если у тебя реально возвращается массив, то нужно будет поправить.
    // Если возвращается один объект, а ты хочешь список — лучше на бэке сделать список.
    // Для примера предположим, что приходит массив задач.
    const data = await apiRequest("/tasks/");

    let tasks = [];
    if (Array.isArray(data)) {
      tasks = data;
    } else if (data && typeof data === "object") {
      // Если бэк возвращает один объект, а не массив — просто обернём
      tasks = [data];
    }

    renderTasks(tasks);
  } catch (err) {
    tasksListEl.innerHTML = "Ошибка загрузки задач: " + err.message;
  }
}

function renderTasks(tasks) {
  tasksListEl.innerHTML = "";

  if (tasks.length === 0) {
    tasksListEl.textContent = "Задач пока нет.";
    return;
  }

  tasks.forEach((task) => {
    const el = document.createElement("div");
    el.className = "task-item" + (task.completed ? " completed" : "");

    const title = document.createElement("div");
    title.className = "task-title";
    title.textContent = task.title || "(без названия)";

    const desc = document.createElement("div");
    desc.className = "task-description";
    desc.textContent = task.description || "";

    const meta = document.createElement("div");
    meta.className = "task-meta";
    meta.textContent =
      "id: " + task.id +
      " | автор: " + task.author_id +
      " | версия: " + task.version +
      (task.completed_at ? " | завершено: " + task.completed_at : "");

    const content = document.createElement("div");
    content.className = "task-content";
    content.appendChild(title);
    if (desc.textContent) content.appendChild(desc);
    content.appendChild(meta);

    const actions = document.createElement("div");
    actions.className = "task-actions";

    // Toggle completed (PATCH)
    const toggleBtn = document.createElement("button");
    toggleBtn.textContent = task.completed ? "Вернуть" : "Готово";
    toggleBtn.addEventListener("click", () => toggleTask(task));

    // Delete
    const deleteBtn = document.createElement("button");
    deleteBtn.textContent = "Удалить";
    deleteBtn.addEventListener("click", () => deleteTask(task));

    actions.appendChild(toggleBtn);
    actions.appendChild(deleteBtn);

    el.appendChild(content);
    el.appendChild(actions);

    tasksListEl.appendChild(el);
  });
}

async function toggleTask(task) {
  const newCompleted = !task.completed;

  // PATCH /tasks/{id}
  // Тело: PatchTaskRequest (у тебя пустой объект в схеме, но логично передавать поля)
  // Отправляем только то, что меняем.
  const body = {
    completed: newCompleted,
  };

  try {
    await apiRequest(`/tasks/${task.id}`, {
      method: "PATCH",
      body: JSON.stringify(body),
    });
    loadTasks();
  } catch (err) {
    alert("Ошибка при обновлении задачи: " + err.message);
  }
}

async function deleteTask(task) {
  if (!confirm("Удалить задачу \"" + (task.title || "без названия") + "\"?")) return;

  try {
    await apiRequest(`/tasks/${task.id}`, {
      method: "DELETE",
    });
    loadTasks();
  } catch (err) {
    alert("Ошибка при удалении задачи: " + err.message);
  }
}

// --- Init ---

loadUserFromStorage();
renderUserState();