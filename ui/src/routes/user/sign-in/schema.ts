import { z } from 'zod';

export const formSchema = z.object({
	email: z.string().email(),
	password: z
		.string()
		.min(8, { message: 'Must be 8 or more characters long' })
		.max(50, { message: 'Must be 50 characters or less' })
});

export type FormSchema = typeof formSchema;
