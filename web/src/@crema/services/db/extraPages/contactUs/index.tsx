export interface ContactUsData {
  sendMessage: SendMessageData;
}

export interface SendMessageData {
  description: string;
}

const contactUsData: ContactUsData = {
  sendMessage: {
    description:
      "The clean and well commented code allows easy customization of the theme. It's designed for describing your app, agency or business.",
  },
};

export default contactUsData;
