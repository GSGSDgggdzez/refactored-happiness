import "./globals.css"
import { ChangeEvent, FormEvent } from "react";
import React from "react"
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Image from "next/image"
export default function Homepage() {


    return (
        <div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-10 mt-9">
                {/* Right Side: Form */}
                <div className="mx-4 md:mx-20 mt-10">
                    <div className="mt-10">
                        <div className="font-bold text-6xl text-green-400">6 Month Free <span className=" text-4xl text-black" >For Our First Users</span> </div>
                        <div className="my-5 font-bold">Enter your credentials below</div>
                    </div>

                    <form >
                        {/* Email field */}
                        <div className="mb-4">
                            <Input
                                name="email"

                                placeholder="Email"
                                type="email"
                                className="w-full border-2 border-gray-300 mt-4"
                            />
                        </div>

                        {/* Password field */}
                        <div className="mb-4">
                            <Input
                                name="password"

                                placeholder="Password"
                                type="password"
                                className="w-full border-2 border-gray-300 mt-4"
                            />
                        </div>

                        {/* Error Message */}
                        {/* {formState.error && <p className="text-red-500 text-sm">{formState.error}</p>} */}

                        {/* Submit Button */}
                        <Button
                            type="submit"

                            className="w-full mt-4 bg-red-500 text-white"
                        >
                            {/* {formState.loading ? "Logging in..." : "Login"} */}
                        </Button>
                    </form>
                </div>
                {/* Right Side: Form */}
                <div className="flex justify-center me-10">
                    <Image
                        src="/side.svg" // Ensure the path is correct
                        alt="Authentication Illustration"
                        width={700} // You can adjust this
                        height={100} // You can adjust this
                        className="max-w-full h-auto" // Make image responsive
                    />
                </div>
            </div>
        </div>
    )
}