import { CartItem } from "@/types";
import { products } from "./products";

export const cartItems:CartItem[] = [
    {
        quantity:1,
        product:products[0]
    },
    {
        quantity:2,
        product:products[1]
    },
    {
        quantity:8,
        product:products[2],
    },
    {
        quantity:5,
        product:products[3]
    }
]