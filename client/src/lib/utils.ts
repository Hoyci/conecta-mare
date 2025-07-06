import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { snakeCase, camelCase } from "lodash-es";

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

      if (Array.isArray(value)) {
        acc[camelKey] = value.map((item) =>
          typeof item === "object" && item !== null ? toCamelCase(item) : item,
        );
      } else if (
        value &&
        typeof value === "object" &&
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

export function formatCentsToBRL(cents: number): string {
  return (cents / 100).toLocaleString("pt-BR", {
    style: "currency",
    currency: "BRL",
  });
}
