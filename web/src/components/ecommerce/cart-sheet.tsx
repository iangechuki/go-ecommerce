"use client"
import { ShoppingCart } from "lucide-react";
import { Button } from "../ui/button";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "../ui/sheet";
import { CartItem } from "./cart-item";
import { ScrollArea } from "../ui/scroll-area";
import { useCart } from "@/lib/stores/cart-store";

export function CartSheet() {
  const {items} =useCart()
  const total = items.reduce((sum, item) => sum + (item.product.price * item.quantity), 0);
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="ghost" size="icon" className="relative">
          <ShoppingCart className="h-5 w-5"/>
          {items.length > 0 && (
            <span className="absolute -top-1 -right-1 bg-primary text-primary-foreground rounded-full h-5 w-5 text-xs flex items-center justify-center">
              {items.length}
            </span>
          )}
        </Button>
      </SheetTrigger>
      <SheetContent className="sm:max-w-lg">
        <SheetHeader>
          <SheetTitle>Shopping Cart ({items.length})</SheetTitle>
        </SheetHeader>
        <ScrollArea className="h-[calc(100vh-140px)] px-0">
          <div className="space-y-2 py-2">
            {items.map((item) => (
              <CartItem key={item.product.id} item={item} />
            ))}
          </div>
        </ScrollArea>

        <div className="border-t pt-2">
          <div className="flex justify-between font-semibold text-lg">
            <span>Total:</span>
            <span>${total.toFixed(2)}</span>
          </div>
          <Button className="w-full mt-2">Checkout</Button>
        </div>
      </SheetContent>
    </Sheet>
  );
}
