import React from 'react'
import './globals.css';
import Image from 'next/image';


export default function RootLayout({

    children,
}: {
    children: React.ReactNode
}) {
    return (
        <html>
            <body className=''>
                <div>
                    <nav className="flex justify-between items-center bg-white text-green-400 px-5 lg:px-10 mt-2">
                        <div className="logo flex items-center w-full space-x-4 transition-all duration-300 ease-in-out hover:scale-105 hover:text-green-500">
                            <Image src="/icon.svg" alt="Logo" width={50} height={100} className="transition-all duration-300 transform hover:scale-110" />
                            <div className="text-4xl font-bold text-green-400 transition-all duration-300 transform hover:text-green-600">
                                SaasDevkit
                            </div>
                        </div>

                        <div className="buttons flex justify-end gap-3">
                            <div>
                                <button className="border-2 border-green-400 text-green-400 font-bold py-2 px-5 mt-2 rounded-3xl flex items-center gap-2 hover:bg-green-100 hover:border-green-600 hover:text-green-600 transition duration-100 ease-in-out">
                                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="size-5">
                                        <path strokeLinecap="round" strokeLinejoin="round" d="M7.5 8.25h9m-9 3H12m-9.75 1.51c0 1.6 1.123 2.994 2.707 3.227 1.129.166 2.27.293 3.423.379.35.026.67.21.865.501L12 21l2.755-4.133a1.14 1.14 0 0 1 .865-.501 48.172 48.172 0 0 0 3.423-.379c1.584-.233 2.707-1.626 2.707-3.228V6.741c0-1.602-1.123-2.995-2.707-3.228A48.394 48.394 0 0 0 12 3c-2.392 0-4.744.175-7.043.513C3.373 3.746 2.25 5.14 2.25 6.741v6.018Z" />
                                    </svg>
                                    About
                                </button>
                            </div>
                            <div>
                                <button className="overflow-hidden relative w-32 p-2  bg-black text-white border-none rounded-md text-xl font-bold cursor-pointer z-10 group mt-2">
                                    Join Beta
                                    <span className="absolute w-36 h-32 -top-8 -left-2 bg-green-200 rounded-full transform scale-x-0 group-hover:scale-x-100 transition-transform group-hover:duration-500 duration-1000 origin-bottom" />
                                    <span className="absolute w-36 h-32 -top-8 -left-2 bg-green-400 rounded-full transform scale-x-0 group-hover:scale-x-100 transition-transform group-hover:duration-700 duration-700 origin-bottom" />
                                    <span className="absolute w-36 h-32 -top-8 -left-2 bg-green-600 rounded-full transform scale-x-0 group-hover:scale-x-100 transition-transform group-hover:duration-1000 duration-500 origin-bottom" />
                                    <span className="group-hover:opacity-100 group-hover:duration-1000 duration-100 opacity-0 absolute top-2.5 left-6 z-10">Welcome!</span>
                                </button>

                            </div>
                        </div>
                    </nav>
                </div>
                {children}
            </body>
        </html>
    )
}