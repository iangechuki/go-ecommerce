"use client";

import { RegisterForm } from "@/components/auth/register-form";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import Link from "next/link";

export default function RegisterPage(){
    return (
        <div className="flex min-h-screen items-center justify-center">
        <Card className="w-full max-w-md">
            <CardHeader>
                <CardTitle className="text-2xl">Sign Up</CardTitle>
            </CardHeader>
            <CardContent>
                <RegisterForm/>
                <div className="mt-4 text-center text-sm">
                    Already have an account?{" "}
                    <Link href="/login" className="underline">
                    Sign In
                    </Link>
                </div>
            </CardContent>
        </Card>
    </div>
    )

}