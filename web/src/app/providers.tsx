"use client";

import { AuthProvider } from "@/context/auth-context";
import { ThemeProvider as NextThemesProvider } from "next-themes";


export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <NextThemesProvider
      attribute="class"
      defaultTheme="system"
      enableSystem
      disableTransitionOnChange
    >
      <AuthProvider>
      {children}
      </AuthProvider>
    </NextThemesProvider>
  );
}