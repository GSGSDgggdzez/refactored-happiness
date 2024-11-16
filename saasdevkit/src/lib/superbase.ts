import { createClient } from '@supabase/supabase-js'

const supabaseUrl = 'https://flldpiweduokkatledat.supabase.co'
const supabaseKey = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZsbGRwaXdlZHVva2thdGxlZGF0Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MzE1ODExNTUsImV4cCI6MjA0NzE1NzE1NX0.t4qvnRjrx1MsF5uEm6JgT9_p5Feq1rBvOza7CRiM7C4'

export const supabase = createClient(supabaseUrl, supabaseKey)
