import { Product } from "@/types";

// constants/products.ts
export const products:Product[] = [
    {
      id: "1",
      name: "Premium Wireless Headphones",
      price: 299.99,
      description: "Noise-canceling Bluetooth headphones with 40hr battery",
      category: "electronics",
      images: [
        "/images/headphones-1.jpg",
        "/images/headphones-2.jpg"
      ],
      stock: 15,
      rating: 4.8,
      reviews: 142,
      originalPrice:200,
    },
    {
      id: "2",
      name: "Smartphone Pro Max",
      price: 999.99,
      description: "Latest generation smartphone with edge-to-edge display and advanced camera",
      category: "electronics",
      images: [
        "/images/smartphone-1.jpg",
        "/images/smartphone-2.jpg"
      ],
      stock: 20,
      rating: 4.6,
      reviews: 230
    },
    {
      id: "3",
      name: "Ultra HD 4K Television",
      price: 1499.99,
      description: "65-inch Ultra HD 4K Smart TV with HDR and voice control",
      category: "electronics",
      images: [
        "/images/tv-1.jpg",
        "/images/tv-2.jpg"
      ],
      stock: 10,
      rating: 4.7,
      reviews: 180
    },
    {
      id: "4",
      name: "Noise-Canceling Earbuds",
      price: 149.99,
      description: "Compact earbuds with active noise cancellation and superior sound quality",
      category: "electronics",
      images: [
        "/images/earbuds-1.jpg",
        "/images/earbuds-2.jpg"
      ],
      stock: 30,
      rating: 4.5,
      reviews: 95
    },
    {
      id: "5",
      name: "Designer Leather Jacket",
      price: 349.99,
      description: "Stylish leather jacket crafted with premium materials",
      category: "fashion",
      images: [
        "/images/jacket-1.jpg",
        "/images/jacket-2.jpg"
      ],
      stock: 8,
      rating: 4.8,
      reviews: 67
    },
    {
      id: "6",
      name: "Elegant Wristwatch",
      price: 199.99,
      description: "Classic analog wristwatch with stainless steel band",
      category: "fashion",
      images: [
        "/images/watch-1.jpg",
        "/images/watch-2.jpg"
      ],
      stock: 25,
      rating: 4.3,
      reviews: 120
    },
    {
      id: "7",
      name: "Modern Sofa",
      price: 899.99,
      description: "Comfortable three-seater sofa with contemporary design",
      category: "home",
      images: [
        "/images/sofa-1.jpg",
        "/images/sofa-2.jpg"
      ],
      stock: 5,
      rating: 4.4,
      reviews: 45
    },
    {
      id: "8",
      name: "Stainless Steel Cookware Set",
      price: 249.99,
      description: "10-piece non-stick cookware set perfect for modern kitchens",
      category: "home",
      images: [
        "/images/cookware-1.jpg",
        "/images/cookware-2.jpg"
      ],
      stock: 12,
      rating: 4.6,
      reviews: 80
    },
    {
      id: "9",
      name: "Fitness Tracker",
      price: 99.99,
      description: "Wearable fitness tracker with heart rate monitor and sleep tracking",
      category: "electronics",
      images: [
        "/images/fitness-1.jpg",
        "/images/fitness-2.jpg"
      ],
      stock: 50,
      rating: 4.2,
      reviews: 210
    },
    {
      id: "10",
      name: "Casual Sneakers",
      price: 79.99,
      description: "Comfortable and stylish sneakers suitable for everyday wear",
      category: "fashion",
      images: [
        "/images/sneakers-1.jpg",
        "/images/sneakers-2.jpg"
      ],
      stock: 40,
      rating: 4.5,
      reviews: 150
    },
    {
      id: "11",
      name: "Portable Bluetooth Speaker",
      price: 129.99,
      description: "Compact speaker with powerful sound and water-resistant design",
      category: "electronics",
      images: [
        "/images/speaker-1.jpg",
        "/images/speaker-2.jpg"
      ],
      stock: 35,
      rating: 4.7,
      reviews: 210
    }
  ];
  