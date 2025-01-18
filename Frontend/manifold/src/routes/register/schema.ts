import { z } from "zod"

export const registerSchema = z.object({
  first_name: z.string().min(2).max(20),
  last_name: z.string().min(2).max(20),
  company: z.string().nullable(),
  phone: z.string().nullable(), 
  email: z.string().email(),
  password: z.string().min(5),
  re_password: z.string().min(5),
  parent_user: z.string().nullable(),
  roles: z.array(z.number()).default([]),
  plan: z.number().default(1)
})

export type RegisterSchema = typeof registerSchema;
