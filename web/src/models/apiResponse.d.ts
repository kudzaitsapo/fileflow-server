import { JsonError } from "./error";

export interface Meta {
  total_records: number;
  limit: number;
  offset: number;
}

export interface JsonResponse<T> {
  success: boolean;
  result: T;
  error: JsonError;
  meta: Meta;
}
