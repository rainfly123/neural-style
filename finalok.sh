# To run this script you'll need to download the ultra-high res
# scan of Starry Night from the Google Art Project, available here:
# https://commons.wikimedia.org/wiki/File:Van_Gogh_-_Starry_Night_-_Google_Art_Project.jpg

STYLE_IMAGE=examples/inputs/starry_night.jpg
CONTENT_IMAGE=examples/inputs/brad_pitt.jpg

STYLE_WEIGHT=5e2
STYLE_SCALE=1.0

th neural_style.lua \
  -content_image $CONTENT_IMAGE \
  -style_image $STYLE_IMAGE \
  -init image -init_image ot3.png \
  -style_scale $STYLE_SCALE \
  -print_iter 1 \
  -style_weight $STYLE_WEIGHT \
  -image_size 3500 \
  -num_iterations 150 \
  -output_image ok.png \
  -tv_weight 0 \
  -lbfgs_num_correction 5 \
  -gpu 0,1 \
  -multigpu_strategy 8 \
  -backend cudnn -cudnn_autotune \
  -model_file models/nin_imagenet_conv.caffemodel \
  -proto_file models/train_val.prototxt  \
  -content_layers relu0,relu3,relu7,relu12 \
  -style_layers relu0,relu3,relu7,relu12
