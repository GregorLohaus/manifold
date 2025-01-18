<script lang="ts">
  import * as Form from "$lib/components/ui/form"
  import { Input } from "$lib/components/ui/input"
  import * as Alert from "$lib/components/ui/alert/index.js";
  import { registerSchema } from "./schema";
  import { superForm } from "sveltekit-superforms";
	import { zodClient } from "sveltekit-superforms/adapters";
  export let data;
  const form = superForm(data, {
    validators: zodClient(registerSchema)
  })
  const { form: formData, message: msg,enhance } = form;
</script>
<form method="POST" use:enhance>
  <Form.Field {form} name="first_name">
    <Form.Control let:attrs>
      <Form.Label> Firstname* </Form.Label>
      <Input { ...attrs } bind:value={$formData.first_name} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

  <Form.Field {form} name="last_name">
    <Form.Control let:attrs>
      <Form.Label> Lastname* </Form.Label>
      <Input { ...attrs } bind:value={$formData.last_name} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

  <Form.Field {form} name="company">
    <Form.Control let:attrs>
      <Form.Label> Company </Form.Label>
      <Input { ...attrs } bind:value={$formData.company} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

  <Form.Field {form} name="phone">
    <Form.Control let:attrs>
      <Form.Label> Phone </Form.Label>
      <Input { ...attrs } bind:value={$formData.phone} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

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

  <Form.Field {form} name="re_password">
    <Form.Control let:attrs>
      <Form.Label> Repeat Password* </Form.Label>
      <Input type="password" { ...attrs } bind:value={$formData.re_password} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>

  <Form.Field {form} name="parent_user">
    <Form.Control let:attrs>
      <Form.Label> Mainuser </Form.Label>
      <Input { ...attrs } bind:value={$formData.parent_user} />
    </Form.Control>
    <Form.FieldErrors/>
  </Form.Field>
  {#if $msg}
    <Alert.Root>
      <Alert.Title>
        Oops and Error accured :(
      </Alert.Title>
        <Alert.Description>
          {#if $msg == 101}
            User already exists, continue to <a href="/verify"> verification </a>
          {:else}
            {$msg}
          {/if}
        </Alert.Description>
    </Alert.Root>
  {/if}
  <Form.Button class="w-full mt-3">Register</Form.Button>
</form>
