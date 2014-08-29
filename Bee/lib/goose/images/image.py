# -*- coding: utf-8 -*-


class Image(object):

    def __init__(self):
        # holds the Element node of the image we think is top dog
        self.top_image_node = None

        # holds the src of the image
        self.src = ""

        # how confident are we in this image extraction?
        # the most images generally the less confident
        self.confidence_score = float(0.0)

        # Height of the image in pixels
        self.height = 0

        # width of the image in pixels
        self.width = 0

        # what kind of image extraction was used for this?
        # bestGuess, linkTag, openGraph tags?
        self.extraction_type = "NA"

        # stores how many bytes this image is.
        self.bytes = long(0)

    def get_src(self):
        return self.src


class ImageDetails(object):

    def __init__(self):

        # the width of the image
        self.width = 0

        # height of the image
        self.height = 0

        # the mime_type of the image JPEG / PNG
        self.mime_type = None

    def get_width(self):
        return self.width

    def set_width(self, width):
        self.width = width

    def get_height(self):
        return self.height

    def set_height(self, height):
        self.height = height

    def get_mime_type(self):
        return self.mime_type

    def set_mime_type(self, mime_type):
        self.mime_type = mime_type


class LocallyStoredImage(object):

    def __init__(self, src='', local_filename='',
        link_hash='', bytes=long(0), file_extension='', height=0, width=0):
        self.src = src
        self.local_filename = local_filename
        self.link_hash = link_hash
        self.bytes = bytes
        self.file_extension = file_extension
        self.height = height
        self.width = width
