"use client"

import { ForgotPasswordForm } from "@/components/auth/forgot-password-form"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import Link from "next/link"

export default function ForgotPasswordPage(){
    return (
        <div className="flex min-h-screen items-center justify-center">
            <Card className="w-full max-w-md">
                <CardHeader>
                    <CardTitle className="text-2xl">
                        Forgot Password
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    <ForgotPasswordForm/>
                    <div className="mt-4 text-center text-sm">
                        Remember your password?{" "}
                        <Link href="/login" className="underline">
                            Sign In</Link>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}