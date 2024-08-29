export type ServerResponse<T = any> = {
  message: string;
  data: T;
};

export type Todo = {
  id: string;
  text: string;
  isDone: boolean;
};

export type CreateTodoDto = {
  text: string;
};

export type UpdateTodoDto = {
  text: string;
  isDone: boolean;
};
