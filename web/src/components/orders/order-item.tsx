import { Order, OrderStatus } from "@/types";
import { Badge } from "../ui/badge";
import { format } from "date-fns";
import { Button } from "../ui/button";
import Link from "next/link";

interface OrderItemParams {
  order: Order;
}

export function OrderItem({ order }: OrderItemParams) {
  const getBadgeVariant = (status: OrderStatus): "destructive" | "secondary" | "default" | "outline" => {
    switch (status) {
      case OrderStatus.DELIVERED:
        return "secondary"; // Adjusted to a valid variant
      case OrderStatus.PENDING:
        return "default"; // Adjusted to a valid variant
      case OrderStatus.CANCELLED:
        return "destructive"; // This one is already valid
      case OrderStatus.PROCESSING:
        return "outline"; // Adjusted to a valid variant
      default:
        return "secondary"; // Fallback
    }
  };

  return (
    <div className="border rounded-lg p-4">
      <div className="flex justify-between items-start mb-2">
        <div>
          <h3 className="font-medium">Order #{order.id}</h3>
          <p className="text-sm text-muted-foreground">{format(new Date(order.date), "MMM dd, yyyy")}</p>
        </div>
        <Badge variant={getBadgeVariant(order.status)}>{order.status}</Badge>
      </div>
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-4">
        <div>
          <p className="text-sm text-muted-foreground">Total</p>
          <p className="font-medium">${order.total.toFixed(2)}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">Items</p>
          <p className="font-medium">{order.items.length}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">Payment</p>
          <p className="font-medium">{order.paymentMethod}</p>
        </div>
        <div>
          <p className="text-sm text-muted-foreground">Tracking</p>
          <p className="font-medium">{order.trackingNumber || "N/A"}</p>
        </div>
      </div>
      <Button variant="outline" size="sm" asChild>
        <Link href={`/orders/${order.id}`}>View Details</Link>
      </Button>
    </div>
  );
}
