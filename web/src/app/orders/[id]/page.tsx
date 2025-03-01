import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { orders } from "@/constants/orders";
import { products } from "@/constants/products";
import { ArrowLeft, Printer } from "lucide-react";
import Image from "next/image";
import Link from "next/link";

export default function Page({ params }: { params: { id: string } }) {
  const order = orders.find((order) => order.id === params.id);

  if (!order) {
    return (
      <div className="container py-8 text-center">
        <h1 className="text-2xl font-bold mb-4">Order not found</h1>
        <Button asChild>
          <Link href="/orders">View all orders</Link>
        </Button>
      </div>
    );
  }
  const getProduct = (productId: string) => 
    products.find(p => p.id === productId);

  const subtotal = order.items.reduce((sum, item) => sum + (item.price * item.quantity), 0);
  const shipping = subtotal > 100 ? 0 : 15;
  const tax = subtotal * 0.07;
  return (
    <div className="container py-8">
      <div className="max-w-4xl mx-auto ">
        <div className="flex justify-between items-start mb-6">
          <Button variant={"ghost"} asChild>
            <Link href="/orders">
              <ArrowLeft className="mr-2 h-4 w-4" />
              Back to Orders
            </Link>
          </Button>
          <Button variant={"outline"}>
            <Printer className="mr-2 h-4 w-4" />
            Print
          </Button>
        </div>
        <div className="bg-card rounded-lg p-6 shadow-sm">
          <div className="flex justify-between items-start mb-8">
            <div>
              <h1>Order #{order.id}</h1>
              <p>{order.date}</p>
            </div>
            <Badge
              variant={
                order.status === "delivered"
                  ? "default"
                  : order.status === "processing"
                  ? "secondary"
                  : order.status === "cancelled"
                  ? "destructive"
                  : "outline"
              }
            >
              {order.status}
            </Badge>
          </div>
          <div className="grid md:grid-cols-2 gap-8 mb-8">
            <div>
              <h2 className="font-semibold mb-2">Shipping Address</h2>
              <div className="space-y-1 text-sm text-muted-foreground">
                <p>John Doe</p>
                <p>123 Main Street</p>
                <p>New York, NY 10001</p>
                <p>United States</p>
              </div>
            </div>

            <div>
              <h2 className="font-semibold mb-2">Order Details</h2>
              <div className="space-y-1 text-sm text-muted-foreground">
                <p>Payment Method: {order.paymentMethod}</p>
                {order.trackingNumber && (
                  <p>Tracking Number: {order.trackingNumber}</p>
                )}
                <p>Order Total: ${order.total.toFixed(2)}</p>
              </div>
            </div>
          </div>
          <div className="border rounded-lg">
            <div className="p-4 bg-muted/50">
              <h3 className="font-semibold">Order Items</h3>
            </div>
            <div className="p-4">
              {order.items.map((item) => {
                const product = getProduct(item.productId);
                return (
                  <div key={item.productId} className="flex items-center gap-4 py-2">
                    <div className="relative h-16 w-16">
                      {product?.images?.[0] && (
                        <Image
                          src={product.images[0]}
                          alt={product.name}
                          fill
                          className="rounded-md object-cover"
                        />
                      )}
                    </div>
                    <div className="flex-1">
                      <p className="font-medium">{product?.name}</p>
                      <p className="text-sm text-muted-foreground">
                        Quantity: {item.quantity}
                      </p>
                    </div>
                    <div className="text-right">
                      <p className="font-medium">
                        ${(item.price * item.quantity).toFixed(2)}
                      </p>
                      <p className="text-sm text-muted-foreground">
                        ${item.price.toFixed(2)} each
                      </p>
                    </div>
                  </div>
                )
              })}
            </div>
          </div>
          <div className="mt-8 space-y-2 max-w-xs ml-auto">
            <div className="flex justify-between">
              <span>Subtotal:</span>
              <span>${subtotal.toFixed(2)}</span>
            </div>
            <div className="flex justify-between">
              <span>Shipping:</span>
              <span>${shipping.toFixed(2)}</span>
            </div>
            <div className="flex justify-between">
              <span>Tax:</span>
              <span>${tax.toFixed(2)}</span>
            </div>
            <div className="flex justify-between font-bold pt-2 border-t">
              <span>Total:</span>
              <span>${order.total.toFixed(2)}</span>
            </div>
          </div>

        </div>
      </div>
    </div>
  );
}
