"use client";

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { cartItems } from "@/constants/cart-items";
import { useCart } from "@/lib/stores/cart-store";
import { Minus, Plus } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { useState } from "react";

export default function CartPage() {
  const [clientSecret, setClientSecret] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const { items, removeItem, updateQuantity } = useCart();

  const user = false;

  const subtotal = items.reduce(
    (sum, item) => sum + item.product.price * item.quantity,
    0
  );
  const shipping = subtotal > 100 ? 0 : 15;
  const tax = subtotal * 0.07;
  const total = subtotal + shipping + tax;

  const handleCheckout = async () => {
    console.log("cheking out");
  };
  return (
    <div className="container py-8 ml-2 space-y-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">Shopping Cart</h1>
        <Button variant={"outline"} asChild>
          <Link href="/products">Continue shopping</Link>
        </Button>
      </div>
      {items.length === 0 ? (
        <div className="text-center py-12">
          <p className="text-muted-foreground mb-4">Your cart is empty</p>
          <Button asChild>
            <Link href="/products">Browse Products</Link>
          </Button>
        </div>
      ) : (
        <div className="grid md:grid-cols-3 gap-8">
          <div className="md:col-span-2 space-y-4">
            {items.map((item) => (
              <Card key={item.product.id}>
                <CardContent className="flex items-center">
                  <div className="relative aspect-square w-24">
                    <Image
                      src={item.product.images[0]}
                      alt={item.product.name}
                      fill
                      className="rounded-md object-cover"
                    />
                  </div>
                  <div className="flex-1 ml-2">
                    <h3 className="font-medium">{item.product.name}</h3>
                    <p className="text-muted-foreground">
                      ${item.product.price}
                    </p>
                    <div className="flex items-center gap-2 mt-2">
                      <Button
                        variant={"outline"}
                        size="sm"
                        onClick={() =>
                          updateQuantity(item.product.id, item.quantity - 1)
                        }
                        disabled={item.quantity <= 1}
                      >
                        <Minus className="h-4 w-4" />
                      </Button>
                      <Input
                        type="number"
                        value={item.quantity}
                        className="w-20 text-center"
                        onChange={(e) =>
                          updateQuantity(
                            item.product.id,
                            Math.max(1, Number(e.target.value))
                          )
                        }
                        disabled={item.quantity <= 1}
                      />
                      
                      <Button
                        variant={"outline"}
                        size="sm"
                        onClick={() =>
                          updateQuantity(item.product.id, item.quantity + 1)
                        }
                        disabled={item.quantity >= item.product.stock}
                      >
                        <Plus className="h-4 w-4" />
                      </Button>
                    </div>
                  </div>
                  <Button variant={"default"} size="sm" onClick={() => removeItem(item.product.id)}>
                    Remove
                  </Button>
                </CardContent>
              </Card>
            ))}
          </div>
        </div>
      )}
      <div className="md:col-span-1">
        <Card>
          <CardHeader>
            <CardTitle>Order Summary</CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <div className="flex justify-between">
                <span>Subtotal</span>
                <span>${subtotal.toFixed(2)}</span>
              </div>
              <div className="flex justify-between">
                <span>Shipping</span>
                <span>${shipping.toFixed(2)}</span>
              </div>
              <div className="flex justify-between">
                <span>Tax</span>
                <span>${tax.toFixed(2)}</span>
              </div>
              <div className="flex justify-between font-bold border-t pt-2">
                <span>Total</span>
                <span>${total.toFixed(2)}</span>
              </div>
            </div>
            {!user ? (
              <div className="space-y-2">
                <Badge variant="destructive" className="w-full">
                  You must be logged in to checkout
                </Badge>
                <Button className="w-full" asChild>
                  <Link href={`/login?redirect=/cart`}>Login to Checkout</Link>
                </Button>
              </div>
            ) : (
              <Button
                className="w-full"
                onClick={handleCheckout}
                disabled={isLoading}
              >
                {isLoading ? "Processing..." : "Proceed to Checkout"}
              </Button>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
