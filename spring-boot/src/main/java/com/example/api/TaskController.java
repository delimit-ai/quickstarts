package com.example.api;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.ArrayList;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;

@RestController
public class TaskController {

    private final Map<Integer, Task> tasks = new ConcurrentHashMap<>();
    private final AtomicInteger nextId = new AtomicInteger(1);

    @GetMapping("/health")
    public Map<String, String> health() {
        return Map.of("status", "ok");
    }

    @GetMapping("/tasks")
    public ArrayList<Task> listTasks() {
        return new ArrayList<>(tasks.values());
    }

    @PostMapping("/tasks")
    public ResponseEntity<Task> createTask(@RequestBody Map<String, String> body) {
        int id = nextId.getAndIncrement();
        Task task = new Task(id, body.get("title"), body.getOrDefault("description", ""));
        tasks.put(id, task);
        return ResponseEntity.status(HttpStatus.CREATED).body(task);
    }

    @GetMapping("/tasks/{id}")
    public ResponseEntity<Task> getTask(@PathVariable int id) {
        Task task = tasks.get(id);
        if (task == null) return ResponseEntity.notFound().build();
        return ResponseEntity.ok(task);
    }

    @DeleteMapping("/tasks/{id}")
    public ResponseEntity<Void> deleteTask(@PathVariable int id) {
        if (tasks.remove(id) == null) return ResponseEntity.notFound().build();
        return ResponseEntity.noContent().build();
    }
}
