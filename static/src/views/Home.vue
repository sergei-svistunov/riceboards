<template>
  <MDBContainer>
    <MDBRow>
      <MDBCol>
        <h1>What is the RICE scoring model for prioritization?</h1>
        <p>The RICE scoring model is a prioritization framework designed to help product managers determine which
          products, features, and other initiatives to put on their roadmaps by scoring these items according to four
          factors. These factors, which form the acronym RICE, are reach, impact, confidence, and effort.</p>
        <p>Using a scoring model such as RICE can offer product teams a three-fold benefit. First, it can enable product
          managers to make better-informed decisions, minimize personal biases in decision making, and help them defend
          their priorities to other stakeholders such as the executive staff.</p>

        <h1>How does the RICE scoring model work?</h1>
        <p>To use the RICE scoring model, you evaluate each of your competing ideas (new products, product extensions,
          features, etc.) by scoring them according to the following formula: </p>
      </MDBCol>
    </MDBRow>
    <MDBRow>
      <MDBCol md="6" offset="3">
        <img src="../assets/rice.svg" class="img-fluid w-100" alt="rice diagram">
      </MDBCol>
    </MDBRow>
    <MDBRow>
      <MDBCol>
        <h2>Reach</h2>
        <p>The first factor in determining your RICE score is to get a sense of how many people you estimate your
          initiative will reach in a given timeframe.</p>
        <p>You have to decide both what “reach” means in this context and the timeframe over which you want to measure
          it. You can choose any time period—one month, a quarter, etc.—and you can decide that reach will refer to the
          number of customer transactions, free-trial signups, or how many existing users try your new feature.</p>
        <p>Your reach score will be the number you’ve estimated. For example, if you expect your project will lead to
          150 new customers within the next quarter, your reach score is 150. On the other hand, if you estimate your
          project will deliver 1,200 new prospects to your trial-download page within the next month, and that 30% of
          those prospects will sign up, your reach score is 360.</p>

        <h2>Impact</h2>
        <p>Impact can reflect a quantitative goal, such as how many new conversions for your project will result in when
          users encounter it, or a more qualitative objective such as increasing customer delight.</p>
        <p>Even when using a quantitative metric (“How many people who see this feature will buy the product?”),
          measuring impact will be difficult, because you won’t necessarily be able to isolate your new project as the
          primary reason (or even a reason at all) for why your users take action. If measuring the impact of a project
          after you’ve collected the data will be difficult, you can assume that estimating it beforehand will also be a
          challenge.</p>

        <h2>Confidence</h2>
        <p>The confidence component of your RICE score helps you control for projects in which your team has data to
          support one factor of your score but is relying more on intuition for another factor.</p>
        <p>For example, if you have data backing up your reach estimate but your impact score represents more of a gut
          feeling or anecdotal evidence, your confidence score will help account for this.</p>

        <h2>Effort</h2>
        <p>We have discussed alll of the factors to this point—reach, impact, confidence—represent the numerators in the
          RICE scoring equation. Effort represents the denominator.</p>
        <p>In other words, if you think of RICE as a cost-benefit analysis, the other three components are all potential
          benefits while effort is the single score that represents the costs.</p>
        <p>Quantifying effort in this model is similar to scoring reach. You simply estimate the total number of
          resources (product, design, engineering, testing, etc.) needed to complete the initiative over a given period
          of time—typically “person-months”—and that is your score.</p>
      </MDBCol>
    </MDBRow>
  </MDBContainer>
</template>

<script lang="ts">
import {defineComponent} from 'vue'
import {MDBCol, MDBContainer, MDBRow} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'Home',
  components: {MDBContainer, MDBRow, MDBCol},
});
</script>
