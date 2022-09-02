resource "aws_mq_configuration" "broker" {
  description    = "the best config"
  name           = "test1"
  engine_type    = "ActiveMQ" //RABBITMQ
  engine_version = "5.15.9"

  data = <<DATA
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<broker xmlns="http://activemq.apache.org/schema/core">
  <plugins>
    <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
    <statisticsBrokerPlugin/>
    <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
  </plugins>
</broker>
DATA

  tags = { name = "some_tags" }

}
