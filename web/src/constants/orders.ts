// constants/orders.ts
import { OrderStatus, Order } from "@/types";
import { products } from "./products";

export const orders: Order[] = [
  {
    id: "ORD-001",
    date: "2024-03-15",
    total: 599.98,
    status: OrderStatus.DELIVERED, // ✅ Using the enum
    items: [
      {
        productId: "1",
        quantity: 2,
        price: products[0].price
      },
      {
        productId: "10",
        quantity: 1,
        price: products[9].price
      }
    ],
    paymentMethod: "credit_card",
    trackingNumber: "UPS-134567890"
  },
  {
    id: "ORD-002",
    date: "2024-03-12",
    total: 1749.98,
    status: OrderStatus.PROCESSING, // ✅ Using the enum
    items: [
      {
        productId: "3",
        quantity: 1,
        price: products[2].price
      },
      {
        productId: "8",
        quantity: 1,
        price: products[7].price
      }
    ],
    paymentMethod: "paypal",
    trackingNumber: "FEDEX-987654321"
  },
  {
    id: "ORD-003",
    date: "2024-03-10",
    total: 429.98,
    status: OrderStatus.PENDING, // ✅ Using the enum
    items: [
      {
        productId: "5",
        quantity: 1,
        price: products[4].price
      },
      {
        productId: "6",
        quantity: 1,
        price: products[5].price
      }
    ],
    paymentMethod: "credit_card"
  },
  {
    id: "ORD-004",
    date: "2024-03-05",
    total: 129.99,
    status: OrderStatus.CANCELLED, // ✅ Using the enum
    items: [
      {
        productId: "11",
        quantity: 1,
        price: products[10].price
      }
    ],
    paymentMethod: "credit_card"
  }
];
