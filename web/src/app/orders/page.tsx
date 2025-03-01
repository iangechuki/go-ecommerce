import { OrderList } from "@/components/orders/order-list";
import { orders } from "@/constants/orders";

export default function OrdersPage(){
    
    return (
        <div className="container mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-8">Order History</h1>
            <OrderList orders={orders}/>
        </div>
    )
}