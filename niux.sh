#!/bin/bash
# To run this script you'll need to download the ultra-high res
# scan of Starry Night from the Google Art Project, available here:
# https://commons.wikimedia.org/wiki/File:Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg
CONTENT_IMAGE=$1
STYLE_IMAGE=$2

STYLE_WEIGHT=5e2
STYLE_SCALE=1.0

  #-gpu 0,1,2,3 \
  #-multigpu_strategy 2,8,16 \
th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -style_scale $STYLE_SCALE \
  -style_weight $STYLE_WEIGHT \
  -image_size 512 \
  -num_iterations 512\
  -output_image ok.jpg\
  -tv_weight 0 \
  -backend cudnn -cudnn_autotune

cp /style/neural-style-master/ok.jpg /live/www/html/emopic/

