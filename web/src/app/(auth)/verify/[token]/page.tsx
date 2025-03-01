"use client"

import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { useParams, useRouter } from "next/navigation"
import { useState } from "react"

export default function VerifyPage(){
    const router = useRouter()
    const {token} = useParams()
    const [status,setStatus] = useState<string>("Verifying...")
    return (
        <div className="flex min-h-screen items-center justify-center">
            <Card className="w-full max-w-md">
                <CardHeader>
                    <CardTitle className="text-2xl">
                        Account verification
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    <p>{status}</p>
                    <div className="mt-4">
                        <Button onClick={()=>{router.push("/login")}}>Go to Login</Button>
                    </div>
                </CardContent>
            </Card>
        </div>
    )
}