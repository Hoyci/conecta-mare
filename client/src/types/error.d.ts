export {};

declare global {
  interface Error {
    tag?: string;
  }
}
