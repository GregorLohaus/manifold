<script lang="ts">
  import * as Form from "$lib/components/ui/form"
  import { Input } from "$lib/components/ui/input"
  import * as Alert from "$lib/components/ui/alert/index.js";
  import { verifySchema, type VerifySchema } from "./schema";
  import { type SuperValidated, type Infer, superForm } from "sveltekit-superforms";
  import { zodClient } from "sveltekit-superforms/adapters";
  export let data: SuperValidated<Infer<VerifySchema>>;
  console.log(data)
  const form = superForm(data, {
    validators: zodClient(verifySchema)
  })
  const { form: formData, message: msg, enhance } = form;
</script>

<form method="POST" use:enhance>
  <Form.Field {form} name="email">
    <Form.Control let:attrs>
      <Form.Label> Email* </Form.Label>
      <Input { ...attrs } bind:value={$formData.email} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

  <Form.Field {form} name="password">
    <Form.Control let:attrs>
      <Form.Label> Password* </Form.Label>
      <Input type="password" { ...attrs } bind:value={$formData.password} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

  <Form.Field {form} name="registration_key">
    <Form.Control let:attrs>
      <Form.Label> Registrationkey* </Form.Label>
      <Input { ...attrs } bind:value={$formData.registration_key} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>
  {#if $msg}
    <Alert.Root>
      <Alert.Title>
        Oops and Error accured :(
      </Alert.Title>
        <Alert.Description>
          {$msg}
        </Alert.Description>
    </Alert.Root>
  {/if}
  <Form.Button class="w-full mt-3">Verify</Form.Button>
</form>
