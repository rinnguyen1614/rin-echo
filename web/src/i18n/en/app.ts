export const app = {
  my_account: "My account",
  logout: "Logout",
  search: "Search",
  configuration: "Configuration",
  language: "Language",
  settings: "Settings",
  email: "Email",
  company: "Company",
  address: {
    address_line1: "Address Line 1",
    address_line2: "Address Line 2",
    city: "City",
    country: "Country",
    district: "District",
    zipcode: "Zipcode",
    state: "state",
  },
  tree: {
    new_node: "New Node",
    actions: {
      add_child: "Add Child",
      add_root: "Add Root",
    },
  },
  audit_logs: {
    user_info: "User Information",
    action_info: "Action Information",
    other_info: "Other Information",
    fields: {
      username: "Username",
      ip_address: "IP address",
      device_name: "Device name",
      user_agent: "User agent",
      location: "Location",
      operation_name: "Operation name",
      operation_method: "Operation method",
      request_id: "Request ID",
      request_url: "Request URL",
      request_body: "Request body",
      response_body: "Response body",
      error: "Error",
      remark: "Remark",
    },
  },
  menu: {
    admin: "Administrator",
  },
  account: {
    tab: {
      personal_info: "Personal Infor",
      change_password: "Change Password",
      information: "Information",
      notification: "Notification",
    },
    profile: {
      fields: {
        email: "Email",
        full_name: "Full Name",
        update_profile: "Update Profile",
      },
    },
    change_password: {
      fields: {
        current_password: "Current Password",
        new_password: "New Password",
        retype_new_password: "Retype New Password",
      },
      validation: {},
    },
  },
  model: {
    created_at: "Creation Time",
    creator_user_id: "Creator",
    modified_at: "Modifion Time",
    modifier_user_id: "Modifier",
  },
  resources: {
    admin: {
      name: "Administration",
    },
    menus: {
      name: "Menu |||| Menus",
      tabs: {
        resources: "Resources",
        general: "General",
      },
    },
    roles: {
      name: "Role |||| Roles",
      tabs: {
        general: "General",
        permissions: "Permissions",
        menus: "Menus",
      },
      fields: {
        name: "Name",
        slug: "Slug",
        is_default: "Default",
        is_static: "Static",
      },
    },
    resources: {
      name: "Resource |||| Resources",
    },
    users: {
      name: "User |||| Users",
      tabs: {
        general: "General",
        roles: "Roles",
      },
      fields: {
        random_password: "Set random password.",
        change_password_on_next_login: "Should change password on next login",
        send_activation_email: "Send activation email.",
        active: "Active",
        lockout_enabled: "Lockout enabled",
      },
      hepls: {
        username: "",
      },
    },
  },
};
