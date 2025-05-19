resource "aws_bedrockagent_prompt" "pike" {
  name            = "MakePlaylist"
  description     = "My first prompt."
  default_variant = "Variant1"

  customer_encryption_key_arn = "arn:aws:kms:eu-west-2:680235478471:key/2ec48faa-3591-481d-a1cb-fd400a404bf6"

  variant {
    name     = "Variant1"
    model_id = "amazon.titan-text-express-v1"

    inference_configuration {
      text {
        temperature = 0.8
      }
    }

    template_type = "TEXT"
    template_configuration {
      text {
        text = "Make me a {{genre}} playlist consisting of the following number of songs: {{number}}."

        input_variable {
          name = "genre"
        }
        input_variable {
          name = "number"
        }
      }
    }
  }

  tags = {
    pike = "permission"
  }
}
