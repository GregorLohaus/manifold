import { getContext, setContext } from "svelte"
import type { VerifySchema } from "../../routes/verify/schema"
import { writable } from "svelte/store"
import type { Infer } from "sveltekit-superforms";

export function setVerifyFormState(initialData: Infer<VerifySchema> | null) {
  const verifyFormState = writable(initialData);
  setContext('VERIFY_FORM_STATE',verifyFormState);
  return verifyFormState;
}

export function getVerifyFormState() {
  return getContext<Infer<VerifySchema> | null>('VERIFY_FORM_STATE')
}
