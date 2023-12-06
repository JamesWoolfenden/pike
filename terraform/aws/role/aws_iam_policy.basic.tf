resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_location_tracker
          "geo:CreateTracker",
          "geo:DescribeTracker",
          "geo:DeleteTracker",

          //aws_location_route_calculator
          "geo:TagResource",
          "geo:UntagResource",
          "geo:CreateRouteCalculator",
          "geo:DescribeRouteCalculator",
          "geo:ListTagsForResource",
          "geo:DeleteRouteCalculator",
          "geo:UpdateRouteCalculator",

          //aws_location_map
          "geo:TagResource",
          "geo:UntagResource",
          "geo:CreateMap",
          "geo:DescribeMap",
          "geo:DescribeMap",
          "geo:DeleteMap",
          "geo:UpdateMap",

          //aws_location_geofence_collection
          "geo:TagResource",
          "geo:UntagResource",
          "geo:CreateGeofenceCollection",
          "kms:DescribeKey",
          "kms:CreateGrant",
          "geo:DescribeGeofenceCollection",
          "geo:DeleteGeofenceCollection",

          //aws_location_place_index
          "geo:TagResource",
          "geo:CreatePlaceIndex",
          "geo:DescribePlaceIndex",
          "geo:DeletePlaceIndex",
          "geo:UntagResource",
          "geo:UpdatePlaceIndex",

        ],
        "Resource" : "*",
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}
