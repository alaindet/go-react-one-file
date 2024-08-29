import clsx from 'clsx';

import { Todo } from '../../types';
import './todo-list.css';

const CHECKBOX_UNCHECKED_CHARCODE = 9744;
const CHECKBOX_CHECKED_CHARCODE = 9745;
const CHECK_CHARCODE = 10004;
const BACK_LOOP_CHARCODE = 8630;

type TodoListProps = {
  todos: Todo[];
  removingId: Todo['id'] | null;
  onItemToggled: (item: Todo) => void;
  onTryRemovingItem: (item: Todo) => void;
  onCancelRemovingItem: () => void;
  onConfirmRemovingItem: (item: Todo) => void;
};

export function TodoList({
  todos,
  removingId = null,
  onItemToggled,
  onTryRemovingItem,
  onCancelRemovingItem,
  onConfirmRemovingItem,
}: TodoListProps) {
  function onItemRemove(todo: Todo) {
    removingId === todo.id
      ? onConfirmRemovingItem(todo)
      : onTryRemovingItem(todo);
  }

  return (
    <ul className="todo-list">
      {todos.map(todo => (
        <li
          key={todo.id}
          className={clsx({
            todo: true,
            '--done': todo.isDone,
          })}
        >
          <span className="todo__icon">
            {String.fromCharCode(
              todo.isDone
                ? CHECKBOX_CHECKED_CHARCODE
                : CHECKBOX_UNCHECKED_CHARCODE
            )}
          </span>
          <span className="todo__text">{todo.text}</span>
          <span className="todo__actions">
            {removingId === todo.id && (
              <button
                type="button"
                className="todo__remove"
                onClick={onCancelRemovingItem}
              >
                &times;&nbsp;No
              </button>
            )}

            <button
              type="button"
              className="todo__remove"
              onClick={() => onItemRemove(todo)}
            >
              &times;&nbsp;{removingId === todo.id ? 'Remove?' : 'Remove'}
            </button>

            <button
              type="button"
              className="todo__done"
              onClick={() => onItemToggled(todo)}
            >
              {String.fromCharCode(
                todo.isDone ? BACK_LOOP_CHARCODE : CHECK_CHARCODE
              )}
              &nbsp;
              {todo.isDone ? 'Undo' : 'Done'}
            </button>
          </span>
        </li>
      ))}
    </ul>
  );
}
