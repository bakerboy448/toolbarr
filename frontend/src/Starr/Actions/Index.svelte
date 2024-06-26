<script lang="ts">
  export let starrApp: StarrApp
  export let hidden: boolean // Avoid querying backends if hidden.

  import type { Instance, StarrApp } from "/src/libs/config"
  import Applogo from "/src/libs/Applogo.svelte"
  import { conf } from "/src/libs/config"
  import T, { _ } from "/src/libs/Translate.svelte"
  import Action from "./action.svelte"
  import Tabs, { startTab } from "./fragments/tabs.svelte"
  import Fa from "svelte-fa"
  import { faCaretLeft, faCaretRight } from "@fortawesome/free-solid-svg-icons"
  import {
    Accordion,
    AccordionItem,
    Badge,
    Card,
    Col,
    Collapse,
    FormGroup,
    Input,
    InputGroup,
    InputGroupText,
    Row,
  } from "@sveltestrap/sveltestrap"

  let updating = false
  let menuOpen = true
  let showTitle = true
  let tab = startTab

  // Start with default instance.
  let instance = $conf.Instances[starrApp][$conf.Instance[starrApp]]
  $: if (!instance || !$conf.Instances[starrApp].includes(instance)) {
    instance = $conf.Instances[starrApp][$conf.Instance[starrApp]]
  }

  let width: number
  $: small = width < 1200
  $: if (small) menuOpen = true
</script>

<svelte:window bind:innerWidth={width}/>

<Accordion>
  <AccordionItem active on:toggle={()=>showTitle=!showTitle}>

    <!-- This is the page title. It's collapsible. -->
    <span slot="header" style="width:95%;">
      <Applogo style="float:right" size="25px" app={starrApp}/>
      <h4 class="d-inline-block accordian-header">{@html $_("instances."+tab.id)}</h4>
      {#if instance}<Badge color="primary">{$_(instance.Name)}</Badge>{/if}
    </span>
    <p><T id="instances.ToolsSelector" {starrApp}/></p>

     <!-- Instance selector menu. -->
     <FormGroup>
      <InputGroup>
        <InputGroupText class="setting-name">{$_("words.Instance")}</InputGroupText>
        <Input invalid={!instance||!instance.URL} type="select" bind:value={instance}>
        {#if $conf.Instances != undefined && $conf.Instances[starrApp]}
          {#each $conf.Instances[starrApp] as i}
            <option value={i}>{i.Name}: {i.URL}</option>
          {/each}
        {:else}
          <option disabled>- {$_("instances.noInstancesConfigured")} -</option>
        {/if}
        </Input>
      </InputGroup>
    </FormGroup>

    <!-- Display the nav links in the accordion header when the screen is small. -->
    {#if small} <Tabs bind:tab fill pills {starrApp} {updating}/> {/if}
  </AccordionItem>
</Accordion>

<Row>
  <div class="container">
    <Col xs="12">
      {#if !small}
        <!-- this creates a toggler-caret that closes/opens the left side nav link menu -->
        <a href="/" on:click|preventDefault={()=>{menuOpen=!menuOpen}}>
          <Fa class="toggle" pull="left" icon={menuOpen?faCaretRight:faCaretLeft}/>
        </a>
        <!-- Display the nav links in the side bar when the screen is not small. -->
        <div class="left">
          <Collapse horizontal isOpen={menuOpen}>
            <Card color={$conf.Dark?"dark":"secondary"}>
              <Tabs bind:tab showTitle vertical pills {starrApp} {updating}/>
            </Card>
          </Collapse>
        </div>
      {/if}
      <div class="right">
        <!-- Display the selected tool, pass in selected instance. -->
        <Action bind:updating {starrApp} {showTitle} {tab} {hidden} {instance} />
      </div>
    </Col>
  </div>
</Row>

<style>
  .container :global(.toggle) {
    margin-top: -4px;
    margin-left: -6px;
    margin-right: -2px;
  }

  .container {
    display: flex;
  }

  .left {
    margin-top: 1px;
    margin-right: 2px;
    float: left;
    width: max-content;
  }

  .right {
    flex-grow: 1;
  }

  .accordian-header {
    margin-bottom: 0 !important;
  }
</style>
