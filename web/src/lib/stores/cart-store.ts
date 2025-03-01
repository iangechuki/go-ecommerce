import { CartItem, Product } from "@/types";
import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

interface CartState {
  items: CartItem[];
  tempUserId: string | null;
  addItem: (product: Product, quantity: number) => void;
  removeItem: (productId: string) => void;
  updateQuantity: (productId: string, quantity: number) => void;
  clearCart: () => void;
}
export const useCart = create<CartState>()(
  persist(
    (set, get) => ({
      items: [] as CartItem[],
      tempUserId: null,
      addItem: (product: Product, quantity: number) => {
        set((state) => {
          const existingItem = state.items.find(
            (item) => item.product.id === product.id
          );
          if (existingItem) {
            return {
              items: state.items.map((item) =>
                item.product.id === product.id
                  ? { ...item, quantity: item.quantity + quantity }
                  : item
              ),
            };
          }
          return {
            items: [...state.items, { product, quantity }],
          };
        });
      },
      removeItem: (productId: string) => {
        set((state)=>({
            items:state.items.filter((item) => item.product.id !== productId)
        }))
      },
      updateQuantity(productId, quantity) {
        set((state) => ({
          items: state.items.map((item) =>
            item.product.id === productId ? { ...item, quantity } : item
          ),
        }));
      },
      clearCart: () => {
        set(()=>({
            items:[]
        }))
      },
    }),
    {
      name: "cart-storage",
      storage: createJSONStorage(() => localStorage),
      partialize: (state) => ({
        items: state.items,
      }),
    }
  )
);
