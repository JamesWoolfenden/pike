data "aws_glue_script" "pike" {
  language = "PYTHON"

  dag_edge {
    source = "datasource0"
    target = "applymapping1"
  }

  dag_edge {
    source = "applymapping1"
    target = "selectfields2"
  }

  dag_edge {
    source = "selectfields2"
    target = "resolvechoice3"
  }

  dag_edge {
    source = "resolvechoice3"
    target = "datasink4"
  }

  dag_node {
    id        = "datasource0"
    node_type = "DataSource"

    args {
      name  = "database"
      value = "pike"
      //value = "\"${aws_glue_catalog_database.source.name}\""
    }
    #
    #    args {
    #      name  = "table_name"
    #      value = "\"${aws_glue_catalog_table.source.name}\""
    #    }
  }

  dag_node {
    id        = "applymapping1"
    node_type = "ApplyMapping"

    args {
      name  = "mapping"
      value = "[(\"column1\", \"string\", \"column1\", \"string\")]"
    }
  }

  dag_node {
    id        = "selectfields2"
    node_type = "SelectFields"

    args {
      name  = "paths"
      value = "[\"column1\"]"
    }
  }

  dag_node {
    id        = "resolvechoice3"
    node_type = "ResolveChoice"

    args {
      name  = "choice"
      value = "\"MATCH_CATALOG\""
    }

    #    args {
    #      name  = "database"
    #      value = "\"${aws_glue_catalog_database.destination.name}\""
    #    }
    #
    #    args {
    #      name  = "table_name"
    #      value = "\"${aws_glue_catalog_table.destination.name}\""
    #    }
  }

  dag_node {
    id        = "datasink4"
    node_type = "DataSink"

    args {
      name  = "database"
      value = "pike"
      //value = "\"${aws_glue_catalog_database.destination.name}\""
    }
    #
    #    args {
    #      name  = "table_name"
    #      value = "\"${aws_glue_catalog_table.destination.name}\""
    #    }
  }
}
