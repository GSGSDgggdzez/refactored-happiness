import "./globals.css"
import { ChangeEvent, FormEvent } from "react";
import MarketingCards from "@/components/MarketingCards";
import React from "react"
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Image from "next/image"
export default function about () {


    return (
        <div className="">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-10 mx-4 md:mx-20 ">
    {/* Left Side: About Section */}
    <div className="mt-5">
        <div className="mt-5">
            <div className="font-bold text-6xl text-green-400">About Us</div>
            <div className="my-1 text-gray-500">We are dedicated to creating a platform that enhances your experience with a continuous array of innovative features.</div>
            <div className="my-1 text-gray-500">
                Our mission is to empower our users by providing tools that make their lives easier and more enjoyable.
            </div>
        </div>

        <div className="my-4 text-gray-500">
            <p>Join us on this exciting journey as we strive to deliver exceptional value and service. Your feedback is invaluable to us, and we are committed to evolving based on your needs.</p>
        </div>
    </div>
    {/* Right Side: Image */}
    <div className="flex justify-center me-10">
        <Image
            src="/about-illustration.svg" // Ensure the path is correct
            alt="About Us Illustration"
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