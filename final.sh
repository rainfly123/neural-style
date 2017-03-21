#!/bin/bash
# To run this script you'll need to download the ultra-high res
# scan of Starry Night from the Google Art Project, available here:
# https://commons.wikimedia.org/wiki/File:Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg

CONTENT_IMAGE=$1
STYLE_IMAGE=$2

STYLE_WEIGHT=5e2
STYLE_SCALE=1.0

th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -style_scale $STYLE_SCALE \
  -style_weight $STYLE_WEIGHT \
  -image_size 256\
  -output_image out1.png \
  -tv_weight 0 \
  -backend cudnn -cudnn_autotune

th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -init image -init_image out1.png \
  -style_scale $STYLE_SCALE \
  -print_iter 1 \
  -style_weight $STYLE_WEIGHT \
  -image_size 512 \
  -num_iterations 1000 \
  -output_image out2.png \
  -tv_weight 0 \
  -backend cudnn -cudnn_autotune

th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -init image -init_image out2.png \
  -style_scale $STYLE_SCALE \
  -print_iter 1 \
  -style_weight $STYLE_WEIGHT \
  -image_size 1024 \
  -num_iterations 200 \
  -output_image out3.png \
  -tv_weight 0 \
  -backend cudnn -cudnn_autotune

th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -init image -init_image out3.png \
  -style_scale $STYLE_SCALE \
  -print_iter 1 \
  -style_weight $STYLE_WEIGHT \
  -image_size 2048 \
  -num_iterations 100 \
  -output_image out4.png \
  -tv_weight 0 \
  -gpu 0,1 \
  -backend cudnn 

th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -init image -init_image out4.png \
  -style_scale $STYLE_SCALE \
  -print_iter 1 \
  -style_weight $STYLE_WEIGHT \
  -image_size 3620 \
  -num_iterations 50 \
  -save_iter 25 \
  -output_image out5.png \
  -tv_weight 0 \
  -lbfgs_num_correction 5 \
  -gpu 0,1,2,3 \
  -multigpu_strategy 3,6,12 \
  -backend cudnn 


th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -init image -init_image out5.png \
  -style_scale $STYLE_SCALE \
  -print_iter 1 \
  -style_weight $STYLE_WEIGHT \
  -image_size 8000 \
  -num_iterations 200 \
  -output_image ok.jpg\
  -tv_weight 0 \
  -gpu 0,1,2,3 \
  -multigpu_strategy 2,8,16 \
  -backend cudnn \
  -optimizer adam \
  -model_file models/nin_imagenet_conv.caffemodel \
  -proto_file models/train_val.prototxt  \
  -content_layers relu0,relu3,relu7,relu12 \
  -style_layers relu0,relu3,relu7,relu12

rm -rf "*.png"
cp /style/neural-style-master/ok.jpg /live/www/html/emopic/
