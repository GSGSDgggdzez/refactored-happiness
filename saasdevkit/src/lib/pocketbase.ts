// lib/pocketbase.ts
import React, { createContext, useContext, useState, useEffect } from 'react';
import PocketBase from 'pocketbase';

// Initialize PocketBase instance
const pb = new PocketBase('http://127.0.0.1:8090');


export interface loginData {
    email: string;
    name: string;
  }


  export const form = async (userData: loginData): Promise<any> => {
    try {
      // Create the record in a waitlist collection
      const record = await pb.collection('waitlist').create(userData);
      
      console.log('Waitlist entry created successfully:', record);
      return record;
    } catch (error) {
      console.error('Error creating waitlist entry:', error);
      throw error;
    }
  };



