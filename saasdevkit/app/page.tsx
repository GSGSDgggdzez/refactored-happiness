import "./globals.css"
import { ChangeEvent, FormEvent } from "react";
import MarketingCards from "@/components/MarketingCards";
import React from "react"
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Image from "next/image"
export default function Homepage() {


    return (
        <div className="">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-10 mx-4 md:mx-20 ">
                {/* Right Side: Form */}
                <div className="mt-5">
                    <div className="mt-5">
                        <div className="font-bold text-6xl text-green-400">6 Month Free <span className=" text-4xl text-black" >For Our First Users</span> </div>
                        <div className="my-1 text-gray-500">We’ll launch a platform with a continues number of features that you will definitely like.</div>
                        <div className="my-1 text-gray-500">
                            And now you can become the very first user and get a god mode absolutely free.
                        </div>
                    </div>

                    <form >
                        {/* Email field */}
                        <div className="grid grid-cols-2 gap-3" >
                            <div className="">
                                <Input
                                    name="name"
                                    placeholder="Name"
                                    type="text"
                                    className="w-full border-2 border-gray-300 mt-4"
                                />
                            </div>

                            {/* Password field */}
                            <div className="">
                                <Input
                                    name="email"
                                    placeholder="email"
                                    type="email"
                                    className="w-full border-2 border-gray-300 mt-4"
                                />
                            </div>
                        </div>
                        {/* Error Message */}
                        {/* {formState.error && <p className="text-red-500 text-sm">{formState.error}</p>} */}

                        {/* Submit Button */}
                        <Button
                            type="submit"

                            className="w-full mt-4 bg-green-400 text-white"
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
            <div className="mt-10">
                <div>
                <MarketingCards />
                </div>
            </div>
        </div>
    )
}