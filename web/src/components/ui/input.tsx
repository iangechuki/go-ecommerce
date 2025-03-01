import * as React from "react"

import { cn } from "@/lib/utils"
interface InputProps extends React.ComponentProps<"input"> {
  error?: string
}
const Input = React.forwardRef<HTMLInputElement, InputProps>(
  ({error, className, type, ...props }, ref) => {
    return (
      <>
          <input
        type={type}
        className={cn(
          "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          error? "border-red-500 focus-visible:ring-red-500":"border-input focus-visible:ring-ring",
          className
        )}
        ref={ref}
        {...props}
      />
      {error && <p className="text-red-500 text-xs mt-1">{error}</p>}
      </>
  

    )
  }
)
Input.displayName = "Input"

export { Input }
