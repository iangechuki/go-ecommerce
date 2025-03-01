"use client";
import Image from "next/image";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Trash } from "lucide-react";
import { CartItem as CartItemType } from "@/types";
import { useCart } from "@/lib/stores/cart-store";
interface CartItemParams {
  item: CartItemType;
}
export function CartItem({ item }: CartItemParams) {
  const { updateQuantity, removeItem } = useCart();
  return (
    <div className="flex items-center gap-2">
      <div className="relative h-16 w-16">
        <Image
          src={item.product.images[0]}
          alt={item.product.name}
          fill
          className="rounded-md object-cover"
        />
      </div>
      <div className="flex-1 text-sm">
        <h5 className="font-medium">{item.product.name}</h5>
        <p className="text-muted-foreground">${item.product.price}</p>
      </div>
      <div className="flex items-center gap-2">
        <Button
          variant="outline"
          size="icon"
          onClick={() => updateQuantity(item.product.id, item.quantity-1)}
          disabled={item.quantity <= 1}
        >
          -
        </Button>
        <Input
          type="number"
          className="w-16 text-center"
          value={item.quantity}
          onChange={(e) =>
            updateQuantity(item.product.id, Number(e.target.value))
          }
        />
        <Button variant={"outline"} onClick={()=> updateQuantity(item.product.id, item.quantity+1)}>+</Button>
      </div>
      <Button
        variant={"ghost"}
        onClick={() => removeItem(item.product.id)}
      >
        <Trash className="h-4 w-4 text-destructive" />
      </Button>
    </div>
  );
}
