export interface PricingObj {
  id: number;
  description?: string;
  priceDescription?: string;
  tag?: string;
  tagColor?: string;
  title?: string;
  popular?: string;
  price: number;
  priceColor?: string;
  pricingList: PricingList[];
}

export interface PricingList {
  id: number;
  title: string;
}

export interface PricingData {
  pricingOne: PricingObj[];
  pricingFour: PricingObj[];
  pricingTwo: PricingObj[];
}

const pricingData: PricingData = {
  pricingOne: [
    {
      id: 1,
      tag: "Basic",
      tagColor: "#11C15B",
      title: "Basic",
      price: 69,
      pricingList: [
        {
          id: 1,
          title: "All features from previous plan",
        },
        {
          id: 2,
          title: "Memberships and bundles",
        },
        {
          id: 3,
          title: "Advanced quizzes",
        },
        {
          id: 4,
          title: "Private & hidden courses",
        },
        {
          id: 5,
          title: "2 Site admin accounts",
        },
        {
          id: 6,
          title: "5 Course admins/authors",
        },
      ],
    },
    {
      id: 2,
      tag: "Pro",
      tagColor: "#FF8B26",
      title: "Pro",
      price: 349,
      popular: "Chosen by 57% of customers",
      pricingList: [
        {
          id: 1,
          title: "All features from previous plan",
        },
        {
          id: 2,
          title: "Memberships and bundles",
        },
        {
          id: 3,
          title: "Advanced quizzes",
        },
        {
          id: 4,
          title: "Private & hidden courses",
        },
        {
          id: 5,
          title: "2 Site admin accounts",
        },
        {
          id: 6,
          title: "5 Course admins/authors",
        },
      ],
    },
    {
      id: 3,
      tag: "Growth",
      tagColor: "#00905F",
      title: "Growth",
      price: 149,
      pricingList: [
        {
          id: 1,
          title: "All features from previous plan",
        },
        {
          id: 2,
          title: "Memberships and bundles",
        },
        {
          id: 3,
          title: "Advanced quizzes",
        },
        {
          id: 4,
          title: "Private & hidden courses",
        },
        {
          id: 5,
          title: "2 Site admin accounts",
        },
        {
          id: 6,
          title: "5 Course admins/authors",
        },
      ],
    },
  ],
  pricingFour: [
    {
      id: 1,
      priceColor: "#11C15B",
      title: "Basic",
      price: 69,
      pricingList: [
        {
          id: 1,
          title: "All features from previous plan",
        },
        {
          id: 2,
          title: "Memberships and bundles",
        },
        {
          id: 3,
          title: "Advanced quizzes",
        },
        {
          id: 4,
          title: "Private & hidden courses",
        },
        {
          id: 5,
          title: "2 Site admin accounts",
        },
        {
          id: 6,
          title: "5 Course admins/authors",
        },
      ],
    },
    {
      id: 2,
      priceColor: "#FF8B26",
      title: "Pro",
      price: 349,
      pricingList: [
        {
          id: 1,
          title: "All features from previous plan",
        },
        {
          id: 2,
          title: "Memberships and bundles",
        },
        {
          id: 3,
          title: "Advanced quizzes",
        },
        {
          id: 4,
          title: "Private & hidden courses",
        },
        {
          id: 5,
          title: "2 Site admin accounts",
        },
        {
          id: 6,
          title: "5 Course admins/authors",
        },
      ],
    },
    {
      id: 3,
      priceColor: "#00905F",
      title: "Growth",
      price: 149,
      pricingList: [
        {
          id: 1,
          title: "All features from previous plan",
        },
        {
          id: 2,
          title: "Memberships and bundles",
        },
        {
          id: 3,
          title: "Advanced quizzes",
        },
        {
          id: 4,
          title: "Private & hidden courses",
        },
        {
          id: 5,
          title: "2 Site admin accounts",
        },
        {
          id: 6,
          title: "5 Course admins/authors",
        },
      ],
    },
  ],
  pricingTwo: [
    {
      id: 1,
      title: "Free",
      description:
        "Designed to help your building initial community and educational content.",
      price: 19,
      priceColor: "#11C15B",
      priceDescription: "No transaction fees",
      pricingList: [
        {
          id: 1,
          title: "1 course to share privately",
        },
        {
          id: 2,
          title: "No selling option",
        },
        {
          id: 3,
          title: "No Social Marketing",
        },
        {
          id: 4,
          title: "Courses & Pages are not discoverable",
        },
        {
          id: 5,
          title: "No team of helpers",
        },
        {
          id: 6,
          title: "No personal and page Blogs",
        },
      ],
    },
    {
      id: 2,
      title: "Start",
      description:
        "Designed to help your building initial community and educational content.",
      price: 89,
      priceColor: "#FF8B26",
      priceDescription: "10% transaction fees + Stripe fees",
      pricingList: [
        {
          id: 1,
          title: "5 courses",
        },
        {
          id: 2,
          title: "Can sell courses and charge users",
        },
        {
          id: 3,
          title: "Marketing with social media",
        },
        {
          id: 4,
          title: "Courses & Pages are discoverable",
        },
        {
          id: 5,
          title: "No team of helpers",
        },
        {
          id: 6,
          title: "Create your personal and page Blogs",
        },
      ],
    },
    {
      id: 3,
      title: "Scale",
      description:
        "Designed to help your building initial community and educational content.",
      price: 49,
      priceColor: "#00905F",
      priceDescription: "8% transaction fees + Stripe fees",
      pricingList: [
        {
          id: 1,
          title: "Unlimited courses",
        },
        {
          id: 2,
          title: "Can sell courses and charge users",
        },
        {
          id: 3,
          title: "Marketing with social media",
        },
        {
          id: 4,
          title: "Courses & Pages are discoverable",
        },
        {
          id: 5,
          title: "No team of helpers",
        },
        {
          id: 6,
          title: "Create your personal and page Blogs",
        },
      ],
    },
  ],
};

export default pricingData;
