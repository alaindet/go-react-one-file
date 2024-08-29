import { ChangeEvent, KeyboardEvent, useMemo, useState } from 'react';

import './todo-item-form.css';
import clsx from 'clsx';

export type TodoItemFormProps = {
  existingTexts: string[];
  isError: boolean;
  onAdd: (text: string) => void;
};

export function TodoItemForm({
  isError,
  onAdd,
  existingTexts,
}: TodoItemFormProps) {
  const [text, setText] = useState('');

  const existingTextsSet = useMemo(() => {
    return new Set(existingTexts.map(t => t.toLowerCase()));
  }, [existingTexts]);

  const duplicateError = useMemo(() => {
    return isError || existingTextsSet.has(text.toLowerCase());
  }, [existingTextsSet, isError, text]);

  const invalid = useMemo(() => {
    return duplicateError || !text;
  }, [text, duplicateError]);

  function handleInputChange(event: ChangeEvent<HTMLInputElement>) {
    setText(event.target.value);
  }

  function handleClickedAdd() {
    if (!text) {
      return;
    }
    onAdd(text);
    setText('');
  }

  function handleKeyDown(event: KeyboardEvent<HTMLInputElement>) {
    if (event.key !== 'Enter') {
      return;
    }
    if (invalid) {
      return;
    }
    handleClickedAdd();
  }

  return (
    <>
      <div
        className={clsx({
          'todo-item-form': true,
          '--error': duplicateError,
        })}
      >
        <input
          type="text"
          value={text}
          placeholder="Enter a new todo item..."
          onChange={handleInputChange}
          onKeyDown={handleKeyDown}
        />
        <button type="button" onClick={handleClickedAdd} disabled={invalid}>
          Add&nbsp;&#x271A;
        </button>
      </div>
    </>
  );
}
