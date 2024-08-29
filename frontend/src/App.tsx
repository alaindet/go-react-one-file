import { useEffect, useMemo, useState } from 'react';

import { TodoItemForm } from './components/todo-item-form/todo-item-form';
import { CreateTodoDto, ServerResponse, Todo, UpdateTodoDto } from './types';
import './App.css';
import { TodoList } from './components/todo-list/todo-list';
import { extractQueryParams } from './utils/extract-query-params';

export function App() {
  const [port, setPort] = useState('8080');
  const apiUrl = useMemo(() => `http://localhost:${port}/api`, [port]);
  const [removingId, setRemovingId] = useState<Todo['id'] | null>(null);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [duplicateError, setDuplicateError] = useState(false);
  const existingTexts = useMemo(() => todos.map(t => t.text), [todos]);

  useEffect(() => {
    initApiPort();
    fetchTodos();
  }, []);

  function initApiPort() {
    const queryParams = extractQueryParams(window.location.href);
    if (queryParams['apiPort']) {
      setPort(queryParams['apiPort']);
    }
  }

  async function fetchTodos() {
    const url = `${apiUrl}/todos`;
    const rawRes = await fetch(url);
    const res = (await rawRes.json()) as ServerResponse<Todo[]>;
    setTodos(res.data);
  }

  async function handleCreateTodo(text: string) {
    const url = `${apiUrl}/todos`;
    const payload: CreateTodoDto = { text };
    const res = await fetch(url, {
      method: 'POST',
      body: JSON.stringify(payload),
    });

    const isDuplicateErr = res.status === 409;
    setDuplicateError(isDuplicateErr);

    if (isDuplicateErr) {
      return;
    }

    await fetchTodos();
    setDuplicateError(false);
  }

  async function handleItemToggled(item: Todo) {
    const payload: UpdateTodoDto = { text: item.text, isDone: !item.isDone };
    const url = `${apiUrl}/todos/${item.id}`;
    await fetch(url, { method: 'PUT', body: JSON.stringify(payload) });
    await fetchTodos();
  }

  function handleTryRemovingItem(item: Todo) {
    setRemovingId(item.id);
  }

  function handleCancelRemovingItem() {
    setRemovingId(null);
  }

  async function handleConfirmRemovingItem(item: Todo) {
    const url = `${apiUrl}/todos/${item.id}`;
    await fetch(url, { method: 'DELETE' });
    await fetchTodos();
    setRemovingId(null);
  }

  return (
    <>
      <h1>YATA - Yet Another Todo App</h1>
      <TodoItemForm
        onAdd={handleCreateTodo}
        isError={duplicateError}
        existingTexts={existingTexts}
      />

      {todos.length > 0 && (
        <TodoList
          todos={todos}
          removingId={removingId}
          onItemToggled={handleItemToggled}
          onTryRemovingItem={handleTryRemovingItem}
          onCancelRemovingItem={handleCancelRemovingItem}
          onConfirmRemovingItem={handleConfirmRemovingItem}
        />
      )}
    </>
  );
}
