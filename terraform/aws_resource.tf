resource "aws_vpc" "vpc" {
  tags = {
    Description = "Created for ECS cluster default"
    Name        = "ECS default - VPC"
  }
}

resource "aws_s3_bucket" "vue" {
  tags = {
    "xxx" = "ttt"
  }
}

resource "aws_s3_bucket_policy" "vue_bucket_policy" {
  bucket = "first-s3-jackyutajack"
  policy = jsonencode(
    {
      Version = "2012-10-17"
      Statement = [
        {
          Action = "s3:GetObject"
          Effect = "Allow"
          Principal = {
            AWS = "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity EFHWMOPJF1U77"
          }
          Resource = "arn:aws:s3:::first-s3-jackyutajack/*"
          Sid      = "2"
        },
      ]
    }
  )
}

resource "aws_ecr_repository" "ecr" {
  name = "first-ecr-jackyutajack"
}

resource "aws_ecs_cluster" "ecs_cluster" {
  name = "default"
}

resource "aws_alb" "alb" {
}

resource "aws_lb_listener" "port80" {
  load_balancer_arn = "arn:aws:elasticloadbalancing:ap-northeast-1:749794220379:loadbalancer/app/alb-jackyutajack/357313bae4e121a6"
  port              = 80
  default_action {
    type             = "forward"
    target_group_arn = "arn:aws:elasticloadbalancing:ap-northeast-1:749794220379:targetgroup/target-group-ecs/21db75df8a93620f"
  }
}

resource "aws_lb_target_group" "lb_target" {
  port                          = 8080
  protocol                      = "HTTP"
  vpc_id                        = "vpc-07a18b6bc29ec4d38"
  load_balancing_algorithm_type = "round_robin"
  name                          = "target-group-ecs"
  protocol_version              = "HTTP1"
  tags                          = {}
  target_type                   = "ip"
  health_check {
    enabled             = true
    healthy_threshold   = 5
    interval            = 30
    matcher             = "200"
    path                = "/users"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    unhealthy_threshold = 2
  }
  stickiness {
    type    = "lb_cookie"
    enabled = false
  }
}

resource "aws_cloudfront_distribution" "cloudfront" {
  tags            = {}
  is_ipv6_enabled = true
  enabled         = true
  restrictions {
    geo_restriction {
      locations        = []
      restriction_type = "none"
    }
  }
  ordered_cache_behavior {
    allowed_methods = [
      "GET",
      "HEAD",
    ]
    cache_policy_id = "4135ea2d-6df8-44a3-9df3-4b5a84be39ad"
    cached_methods = [
      "GET",
      "HEAD",
    ]
    compress               = true
    default_ttl            = 0
    max_ttl                = 0
    min_ttl                = 0
    path_pattern           = "/users"
    smooth_streaming       = false
    target_origin_id       = "alb-jackyutajack-785790821.ap-northeast-1.elb.amazonaws.com"
    trusted_key_groups     = []
    trusted_signers        = []
    viewer_protocol_policy = "https-only"
  }

  origin {
    connection_attempts = 3
    connection_timeout  = 10
    domain_name         = "first-s3-jackyutajack.s3.ap-northeast-1.amazonaws.com"
    origin_id           = "first-s3-jackyutajack.s3.ap-northeast-1.amazonaws.com"
    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/EFHWMOPJF1U77"
    }
  }

  origin {
    connection_attempts = 3
    connection_timeout  = 10
    domain_name         = "alb-jackyutajack-785790821.ap-northeast-1.elb.amazonaws.com"
    origin_id           = "alb-jackyutajack-785790821.ap-northeast-1.elb.amazonaws.com"
    custom_origin_config {
      http_port                = 80
      https_port               = 443
      origin_keepalive_timeout = 5
      origin_protocol_policy   = "http-only"
      origin_read_timeout      = 30
      origin_ssl_protocols = [
        "TLSv1",
        "TLSv1.1",
        "TLSv1.2",
      ]
    }
  }
  default_cache_behavior {
    cache_policy_id = "658327ea-f89d-4fab-a63d-7e88639e58f6"
    compress        = true
    function_association {
      event_type   = "viewer-request"
      function_arn = "arn:aws:cloudfront::749794220379:function/spa-request-handler"
    }
    viewer_protocol_policy = "redirect-to-https"
    allowed_methods        = ["GET", "HEAD"]
    cached_methods         = ["GET", "HEAD"]
    target_origin_id       = "first-s3-jackyutajack.s3.ap-northeast-1.amazonaws.com"
  }
  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

resource "aws_cloudfront_function" "cloudfront_function" {
  name    = "spa-request-handler"
  runtime = "cloudfront-js-1.0"
  comment = "my function"
  publish = true
  code    = file("${path.module}/src/spa-request-handler.js")
}
