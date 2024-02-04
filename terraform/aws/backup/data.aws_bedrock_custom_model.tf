data "aws_bedrock_custom_model" "pike" {
  provider = aws.central
  model_id = "arn:aws:bedrock:us-east-1:680235478471:custom-model/amazon.titan-text-express-v1:0:8k/ly16hhi765j4"
}
