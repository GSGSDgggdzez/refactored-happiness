"use client";

import "./globals.css"
import { ChangeEvent, FormEvent, useState } from "react";
import MarketingCards from "@/components/MarketingCards";
import React from "react"
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Image from "next/image"
import { form } from "@/lib/pocketbase";

interface form {
    email: string;
    name: string;
    loading: boolean;
    error: string | null;
    isSuccess: boolean;
}
export default function Homepage() {

    const [formState, setFormState] = useState<form>({
        email: "",
        name: "",
        loading: false,
        error: null,
        isSuccess: false
    });

    const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormState((prev) => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();
        setFormState((prev) => ({ ...prev, loading: true, error: null }));

        try {
            const loginData = {
                email: formState.email,
                name: formState.name,
            };
            const user = await form(loginData);
            setFormState(prev => ({ ...prev, isSuccess: true }));
        } catch (err) {
            console.error("Error during user login:", err);
            setFormState((prev) => ({
                ...prev,
                error: "Invalid email or password. Please try again.",
            }));
        } finally {
            setFormState((prev) => ({ ...prev, loading: false }));
        }
    };

    return (
        <div className="min-h-screen flex flex-col">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-10 mx-4 md:mx-20 ">
                <div className="mt-36">
                    <div className="flex-grow ">
                        <div className="font-bold text-6xl text-green-400">6 Month Free <span className=" text-4xl text-black" >For Our First Users</span> </div>
                        {!formState.isSuccess ? (
                            <>
                                <div className="my-1 text-gray-500">We'll launch a platform with a continues number of features that you will definitely like.</div>
                                <div className="my-1 text-gray-500">
                                    And now you can become the very first user and get a god mode absolutely free.
                                </div>
                                <form onSubmit={handleSubmit}>
                                    <div className="grid grid-cols-2 gap-3" >
                                        <div className="">
                                            <Input
                                                name="name"
                                                placeholder="Name"
                                                type="text"
                                                value={formState.name}
                                                onChange={handleInputChange}
                                                className="w-full border-2 border-gray-300 mt-4"
                                            />
                                        </div>
                                        <div className="">
                                            <Input
                                                name="email"
                                                placeholder="email"
                                                value={formState.email}
                                                onChange={handleInputChange}
                                                type="email"
                                                className="w-full border-2 border-gray-300 mt-4"
                                            />
                                        </div>
                                    </div>
                                    {formState.error && <p className="text-red-500 text-sm">{formState.error}</p>}
                                    <Button
                                        type="submit"
                                        disabled={formState.loading}
                                        className="w-full mt-4 bg-green-400 text-white"
                                    >
                                        {formState.loading ? "Logging in..." : "Login"}
                                    </Button>
                                </form>
                            </>
                        ) : (
                            <div className="mt-8 text-2xl text-green-500 font-semibold">
                                Thank you for joining our waiting list! We'll keep you updated on our launch.
                            </div>
                        )}
                    </div>
                </div>
                <div className="flex justify-center me-10">
                    <Image
                        src="/side.svg"
                        alt="Authentication Illustration"
                        width={700}
                        height={100}
                        className="hidden md:block max-w-full h-auto"
                    />
                </div>
            </div>
            <footer className="mt-auto">
                <footer>
                    <MarketingCards />
                </footer>
            </footer>
        </div>

    )
}