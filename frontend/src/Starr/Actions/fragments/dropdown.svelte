<script lang="ts">
  export let form: any
  export let updating: boolean
  export let starrApp: StarrApp
  export let field: string
  export let name = ""

  import T, { _ } from "/src/libs/Translate.svelte"
  import Fa from "svelte-fa"
  import { faCircleInfo } from "@fortawesome/free-solid-svg-icons"
  import { Dropdown, DropdownItem, DropdownMenu, DropdownToggle, Tooltip } from "@sveltestrap/sveltestrap"
  import type { StarrApp } from "/src/libs/config"
  import { onOnce } from "/src/libs/funcs"

  $: title = name ? $_(name) : $_(`instances.${field}Title`)
  let dropdown: HTMLElement

  function toggleAll(key: string, on: boolean) {
    var idx = 0.15 // initial delay to click.
    form.forEach((_: any, i: number) => { // progressively faster.
      onOnce(() => {form[i][key] = on}, idx += 0.08 - (idx/14))
    })
  }
</script>

<Tooltip placement="top" target={dropdown}>
  <T id={`instances.${field}Desc`} {starrApp}/>
</Tooltip>
<Dropdown size="sm">
  <DropdownToggle tag="span">
    <span bind:this={dropdown} class="link">
      {title.match(/\b(\w)/g)?.join('').substring(0,2)} <Fa primaryColor="darkCyan" icon={faCircleInfo}/>
    </span>
  </DropdownToggle>
  <DropdownMenu>
    <DropdownItem header>{title}</DropdownItem>
    <DropdownItem disabled={updating} on:click={() => toggleAll(field, false)}>
      {$_("instances.DisableAll")}
    </DropdownItem>
    <DropdownItem disabled={updating} on:click={() => toggleAll(field, true)}>
      {$_("instances.EnableAll")}
    </DropdownItem>
  </DropdownMenu>
</Dropdown>

<style>
  .link {
    cursor: pointer;
    text-decoration: underline;
    color: rgb(19, 87, 87)
  }
</style>