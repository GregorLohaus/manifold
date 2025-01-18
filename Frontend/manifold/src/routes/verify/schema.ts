import { z } from "zod"

export const verifySchema = z.object({
  email: z.string().email(),
  password: z.string().min(5),
  registration_key: z.string().min(5)
})

export type VerifySchema = typeof verifySchema
