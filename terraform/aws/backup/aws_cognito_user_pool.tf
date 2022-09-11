resource "aws_cognito_user_pool" "pike" {
  name = "pike"

  account_recovery_setting {
    recovery_mechanism {
      name     = "verified_email"
      priority = 2
    }
  }

  admin_create_user_config {
    allow_admin_create_user_only = true
    invite_message_template {
      email_message = "hello {username} here is your password {####}"
      email_subject = "hello"
      sms_message   = "hello {username} here is your password {####}"
    }
  }

  alias_attributes = ["phone_number", "email", "preferred_username"]

  #  auto_verified_attributes = []
  #
  device_configuration {
    challenge_required_on_new_device      = false
    device_only_remembered_on_user_prompt = false
  }

  email_configuration {
    #    configuration_set = ""
    email_sending_account = "COGNITO_DEFAULT"
    #    from_email_address = ""
    #    reply_to_email_address = ""
    #    source_arn = ""
  }

  #  email_verification_message = ""
  #  email_verification_subject = ""
  #
  lambda_config {
    create_auth_challenge = "arn:aws:lambda:eu-west-2:680235478471:function:message-processor"
    #    custom_message = ""
    #    define_auth_challenge = ""
    #    post_authentication = ""
    #    post_confirmation = ""
    #    pre_authentication = ""
    #    pre_sign_up = ""
    #    pre_token_generation = ""
    #    user_migration = ""
    #    verify_auth_challenge_response = ""
    kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
    #    custom_email_sender {
    #      lambda_arn     = ""
    #      lambda_version = ""
    #    }
    #    custom_sms_sender {
    #      lambda_arn     = ""
    #      lambda_version = ""
    #    }
  }
  mfa_configuration = "ON"
  password_policy {
    minimum_length                   = 10
    require_lowercase                = true
    require_numbers                  = true
    require_symbols                  = true
    require_uppercase                = true
    temporary_password_validity_days = 7
  }
  schema {
    attribute_data_type      = "String"
    developer_only_attribute = false
    name                     = "pike"
    mutable                  = false
    #    number_attribute_constraints {
    #      max_value = 10
    #      min_value = 1
    #    }
    required = false
    string_attribute_constraints {
      max_length = 10
      min_length = 1
    }

  }


  #  sms_configuration {
  #    external_id    = "9ac53c3a-c421-44da-93dd-b1b6cc8c6bca"
  #    sns_caller_arn = "arn:aws:iam::680235478471:role/service-role/pike-SMS-Role"
  #  }

  sms_verification_message = "Your code is {####}"
  software_token_mfa_configuration {
    enabled = true
  }

  user_pool_add_ons {
    advanced_security_mode = "AUDIT"
  }
  #
  #  //username_attributes = []
  #
  username_configuration {
    case_sensitive = true
  }
  #
  #  verification_message_template {
  #    default_email_option = "CONFIRM_WITH_CODE"
  #    email_message = ""
  #    email_message_by_link = ""
  #    email_subject = ""
  #    email_subject_by_link = ""
  #    sms_message = ""
  #  }
  tags = {
    pike    = "permissions"
    another = "one"
  }
}
