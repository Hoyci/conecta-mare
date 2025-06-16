import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { snakeCase, camelCase } from "lodash";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function toSnakeCase(obj: Record<string, any>): Record<string, any> {
  return Object.entries(obj).reduce(
    (acc, [key, value]) => {
      const snakeKey = snakeCase(key);
      if (
        value &&
        typeof value === "object" &&
        !Array.isArray(value) &&
        !(value instanceof Date)
      ) {
        acc[snakeKey] = toSnakeCase(value);
      } else {
        acc[snakeKey] = value;
      }
      return acc;
    },
    {} as Record<string, any>,
  );
}

export function toCamelCase(obj: Record<string, any>): Record<string, any> {
  return Object.entries(obj).reduce(
    (acc, [key, value]) => {
      const camelKey = camelCase(key);
      if (
        value &&
        typeof value === "object" &&
        !Array.isArray(value) &&
        !(value instanceof Date)
      ) {
        acc[camelKey] = toCamelCase(value);
      } else {
        acc[camelKey] = value;
      }
      return acc;
    },
    {} as Record<string, any>,
  );
}
