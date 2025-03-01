import { Order } from "@/types"
import { Skeleton } from "../ui/skeleton"
import { OrderItem } from "./order-item"

interface OrderListProps {
    orders?: Order[]
}
export function OrderList({ orders=[] }: OrderListProps) {
    return (
        <div className="space-y-4">
        {orders.length > 0 ? (
          orders.map((order) => <OrderItem key={order.id} order={order} />)
        ) : (
          <p className="text-gray-500">No orders found.</p>
        )}
      </div>
    )
}
export function OrderListSkeleton() {
    return (
        <div className="space-y-3">
            {[1, 2, 3].map((i) => (
                <Skeleton key={i} className="h-32 w-full rounded-lg" />
            ))}
        </div>
    )
}