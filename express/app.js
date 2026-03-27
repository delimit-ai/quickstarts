const express = require("express");
const app = express();

app.use(express.json());

const tasks = new Map();
let nextId = 1;

app.get("/health", (_req, res) => {
  res.json({ status: "ok" });
});

app.get("/tasks", (_req, res) => {
  res.json([...tasks.values()]);
});

app.post("/tasks", (req, res) => {
  const { title, description = "" } = req.body;
  const task = { id: nextId++, title, description, done: false };
  tasks.set(task.id, task);
  res.status(201).json(task);
});

app.get("/tasks/:id", (req, res) => {
  const task = tasks.get(Number(req.params.id));
  if (!task) return res.sendStatus(404);
  res.json(task);
});

app.delete("/tasks/:id", (req, res) => {
  const id = Number(req.params.id);
  if (!tasks.has(id)) return res.sendStatus(404);
  tasks.delete(id);
  res.sendStatus(204);
});

const port = process.env.PORT || 3000;
app.listen(port, () => console.log(`TaskForge running on port ${port}`));
