export type Book = {
  id: string;
  title: string;
  genre: string;
  totalPage: number;
  progressPage?: number;
  progressPercentage?: number;
  author: string;
  publisher: string;
  publishedAt: number;
  price: number;
};
