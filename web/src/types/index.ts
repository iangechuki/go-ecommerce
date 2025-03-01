export interface Product {
  id: string;
  name: string;
  price: number;
  description: string;
  category: Category["slug"];
  images: string[];
  stock: number;
  rating: number;
  reviews: number;
  originalPrice?: number;
}

export interface Category {
  name: string;
  slug: string;
}
export interface CartItem {
  product: Product;
  quantity: number;
}
export enum OrderStatus {
  PENDING = "pending",
  PROCESSING = "processing",
  DELIVERED = "delivered",
  CANCELLED = "cancelled",
}
export interface Order {
  id: string;
  date: string;
  total: number;
  status: OrderStatus;
  items: OrderItem[];
  paymentMethod: string;
  trackingNumber?: string;
}

export interface OrderItem {
  productId: string;
  quantity: number;
  price: number;
}

export interface User {
  id: string;
  name: string;
  email: string;
  password: string;
  role: string;
}